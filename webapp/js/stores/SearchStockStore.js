/**
 * Created by yen-chieh on 5/13/16.
 */
import {EventEmitter} from 'events';
import dispatcher from '../dispatcher/appDispatcher.js'
import constants from '../constants/mainConstants.js';

class SearchStockStore extends EventEmitter {
	constructor() {
		super();
		this.quoteData = localStorage.getItem("QuoteData") ? JSON.parse(localStorage.getItem("QuoteData")) : [];
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

		this.emit("change");
	}

	getAllQuoteData(){
		return this.quoteData;
	}

	updateQuoteDataHistory() {
		localStorage.setItem("QuoteData", JSON.stringify(this.quoteData));
	}

	handleAction(action){
		switch(action.type){
			case constants.RECEIVED_QUOTE_DATA:
				this.addQuoteData(action.data);
				this.updateQuoteDataHistory();
				break;
		}
	}

}

const searchStockStore = new SearchStockStore();
dispatcher.register(searchStockStore.handleAction.bind(searchStockStore));

export default searchStockStore;