/**
 * Created by yen-chieh on 5/12/16.
 */

import React from 'react';
import _ from 'underscore';
import SearchStockStore from '../stores/SearchStockStore.js';
import dispatcher from '../dispatcher/appDispatcher.js';
import * as QuoteAction from '../actions/QuoteAction.js';
import QuoteTable from './quote.js';
import constant from '../constants/mainConstants.js'

var SearchPage = React.createClass({
	getInitialState: function () {
		return {
			quoteData: [],
			checkedQuote: []
		}
	},

	componentWillMount: function () {
		SearchStockStore.on('updateQuoteData', this.updateQuoteData);
		SearchStockStore.on('updateCheckedList', this.updateCheckedList);
	},

	componentDidMount: function () {
		var self = this;
		this.$optionButtons = $('button', '#listOption');

		if (this.state.checkedQuote.length == 0) {
			_.each(this.$optionButtons, function (data) {
				$(data).addClass('disabled');
			});
		}

		this.text = $('input[name="symbol"]');
		this.$quoteButton = $('#getQuoteButton');

		this.updateQuoteData();

		setTimeout(function(){
			self.refreshQuotes();
		}, 1000);

		this.refreshQuoteInterval = setInterval(function(){
			self.refreshQuotes();
		}, 25000);

	},

	componentWillUnmount: function(){
		SearchStockStore.removeListener("updateQuoteData", this.updateQuoteData);
		SearchStockStore.removeListener("updateCheckedList", this.updateCheckedList);
		clearInterval(this.refreshQuoteInterval);
	},

	refreshQuotes(){
		if(this.state.quoteData.length == 0){
			return;
		}
		var symbols = "";
		this.state.quoteData.map(function(quote, i){
			symbols += "," + quote.symbol
		});

		QuoteAction.getQuoteData(symbols.substring(1), constant.REFRESH_QUOTE_DATA_SEARCH);
	},

	updateQuoteData: function(){
		this.setState({
			quoteData: SearchStockStore.getAllQuoteData()
		});

		this.$quoteButton.button('reset');
	},

	updateCheckedList: function(){
		this.setState({
			checkedQuote: SearchStockStore.getCheckedQuote()
		});
		if (SearchStockStore.getCheckedQuote().length == 0) {
			_.each(this.$optionButtons, function (data) {
				$(data).addClass('disabled');
			});
		} else {
			_.each(this.$optionButtons, function (data) {
				$(data).removeClass('disabled');
			});
		}
	},

	keyPressed: function (e) {
		if (e.key == "Enter") {
			this.searchQuoteData();
		}
	},

	searchQuoteData: function () {
		this.$quoteButton.button('loading');
		QuoteAction.getQuoteData(this.text.val(), constant.RECEIVED_QUOTE_DATA_SEARCH);
		this.text.val("");
	},

	removeQuote: function () {
		QuoteAction.removeQuoteByCheckedlist();
	},

	addToWatchList: function(e) {
		//QuoteAction.addCheckedQuoteToWatchList();
		if($(e.target).hasClass("disabled")){
			return false;
		}
		this.props.addListButtonClicked();
	},

	render: function () {
		return (
			<div>
				<div id="mainSearchComponent">
					<figure className="highlight">
						<div className="form-inline">
							<div className="form-group">
								<label for="symbolSearchInput" className="sr-only">Enter Stock Here</label>
								<input type="text" name="symbol" className="form-control" id="symbolSearchInput"
											 placeholder="Enter symbol or stock name" onKeyPress={this.keyPressed}/>
							</div>
							<button className="btn btn-sm btn-primary" id="getQuoteButton" onClick={this.searchQuoteData} data-loading-text="Loading..." autocomplete="off">Search</button>
						</div>
					</figure>

					<div id="quoteList">
						<h1>Quote List</h1>

						<div id="listOption">
							<button className="btn btn-sm btn-info" id="watchListButton" onClick={this.addToWatchList}>Add to your watch list</button>
							<button className="btn btn-sm btn-danger" id="deleteButton" onClick={this.removeQuote}>Delete</button>
						</div>

						<QuoteTable quote={this.state.quoteData} checkedQuote={this.state.checkedQuote}/>
					</div>
				</div>
			</div>
		)
	}
});

export default SearchPage;
