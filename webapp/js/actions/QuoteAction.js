/**
 * Created by yen-chieh on 5/13/16.
 */

import dispatcher from '../dispatcher/appDispatcher.js';
import constants from '../constants/mainConstants.js';
import 'whatwg-fetch';

export function getQuoteData(symbol){
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
				type: constants.RECEIVED_QUOTE_DATA,
				data: quoteData
			})
		}
	);
}