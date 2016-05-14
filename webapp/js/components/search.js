/**
 * Created by yen-chieh on 5/12/16.
 */

import React from 'react';
import SearchStockStore from '../stores/SearchStockStore.js';
import dispatcher from '../dispatcher/appDispatcher.js';
import * as QuoteAction from '../actions/QuoteAction.js';
import QuoteTable from './quote.js';

var SearchPage = React.createClass({
	getInitialState: function(){
		return {
			quoteData: SearchStockStore.getAllQuoteData()
		}
	},

	componentWillMount: function(){
		SearchStockStore.on('change', () =>{
			this.setState({
				quoteData: SearchStockStore.getAllQuoteData()
			})
		});
	},

	componentDidMount: function(){
		this.text = $('input[name="symbol"]');
	},

	keyPressed: function(e){
		if (e.key == "Enter") {
			this.searchQuoteData();
		}
	},

	searchQuoteData: function(){
		QuoteAction.getQuoteData(this.text.val());
		this.text.val("");
	},

	render: function(){
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
							<button className="btn btn-sm btn-primary" onClick={this.searchQuoteData}>Search</button>
						</div>
					</figure>

					<div id="quoteList">
						<h1>Quote List</h1>
						<QuoteTable quote={this.state.quoteData} />
					</div>
				</div>
			</div>
		)
	}
});

export default SearchPage;