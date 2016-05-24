/**
 * Created by yen-chieh on 5/13/16.
 */

import dispatcher from '../dispatcher/appDispatcher.js';
import constants from '../constants/mainConstants.js';
import 'whatwg-fetch';

export function getQuoteData(symbol, source){
	dispatcher.dispatch({type: constants.FETCH_DATA_FROM_SERVER});

	fetch(constants.API_DOMAIN + constants.API_SEARCH_STOCK + "?symbol=" + symbol, {
		method: 'GET',
		headers: constants.API_JSON_HEADER
	}).then((response) => response.json())
		.then((data) => {
			if (!data || data.length == 0) {
				return false;
			}

			var quoteData;
			if(data.query.count > 1){
				quoteData = data.query.results.quotes;
			}else{
				quoteData = data.query.results.quote;
			}
			dispatcher.dispatch({
				type: source,
				data: quoteData
			})
		}
	);
}

export function addCheckedQuoteToWatchList(user, checkedList) {
	dispatcher.dispatch({type: constants.FETCH_DATA_FROM_SERVER});

	var data = {
		email: user.email,
		stocks: checkedList
	};

	fetch(constants.API_DOMAIN + constants.API_ADD_WATCH_LIST, {
		method: 'POST',
		headers: constants.API_JSON_HEADER,
		body: JSON.stringify(data)
	}).then((response) => response.json())
		.then((data) => {
			if (!data || data.length == 0) {
				return false;
			}
			dispatcher.dispatch({
				type: constants.WATCH_LIST_UPDATED,
				data: data
			})
		}
	);
}

export function updateCheckedQuoteList(checkedList){
	dispatcher.dispatch({
		type: constants.UPDATE_CHECKED_LIST,
		data: checkedList
	});
}

export function removeQuoteByCheckedlist(){
	dispatcher.dispatch({
		type: constants.REMOVE_QUOTE
	});
}