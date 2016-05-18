/**
 * Created by yen-chieh on 5/13/16.
 */
import {EventEmitter} from 'events';
import dispatcher from '../dispatcher/appDispatcher.js'
import constants from '../constants/mainConstants.js';
import update from 'react-addons-update';

class SearchStockStore extends EventEmitter {
	constructor() {
		super();
		this.quoteData = localStorage.getItem("QuoteData") ? JSON.parse(localStorage.getItem("QuoteData")) : [];
		this.checkedQuote = [];
	}

	addQuoteData(quoteData){
		var self = this;
		if(quoteData.length > 0){
			quoteData.map(function(data, i){
				self.quoteData.unshift(data);
			})
		}else{
			this.quoteData.unshift(quoteData);
		}

		this.emit("updateQuoteData");
	}

	updateCheckedList(checkedList){
		var checkedQuote = this.checkedQuote;

		var index = checkedQuote.indexOf(checkedList.toUpperCase());
		if(index != -1){
			this.checkedQuote = update(checkedQuote, {$splice: [[index, 1]]});
		}else{
			this.checkedQuote = update(checkedQuote, {$push: [checkedList.toUpperCase()]});
		}

		this.emit("updateCheckedList");
	}

	getCheckedQuote() {
		return this.checkedQuote;
	}

	getAllQuoteData(){
		return this.quoteData;
	}

	updateQuoteDataHistory() {
		localStorage.setItem("QuoteData", JSON.stringify(this.quoteData));
	}

	removeQuoteByCheckedList(){
		var self = this;
		var newQuoteData = [];

		this.quoteData.map(function(data,i){
			var index = self.checkedQuote.indexOf(data.symbol.toUpperCase());
			if(index == -1){
				newQuoteData = update(newQuoteData, {$push: [data]})
			}
		});
		this.checkedQuote = [];
		this.quoteData = newQuoteData;
		this.updateQuoteDataHistory();

		this.emit("updateQuoteData");
		this.emit('updateCheckedList');
	}

	handleAction(action){
		switch(action.type){
			case constants.RECEIVED_QUOTE_DATA:
				this.addQuoteData(action.data);
				this.updateQuoteDataHistory();
				break;
			case constants.UPDATE_CHECKED_LIST:
				this.updateCheckedList(action.data);
				break;
			case constants.REMOVE_QUOTE:
				this.removeQuoteByCheckedList();
				break;
		}
	}

}

const searchStockStore = new SearchStockStore();
dispatcher.register(searchStockStore.handleAction.bind(searchStockStore));

export default searchStockStore;