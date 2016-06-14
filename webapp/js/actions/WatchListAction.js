/**
 * Created by yen-chieh on 5/25/16.
 */

import dispatcher from '../dispatcher/appDispatcher.js';
import constants from '../constants/mainConstants.js';
import 'whatwg-fetch';

export function getWatchList(email){
	dispatcher.dispatch({type: constants.FETCH_DATA_FROM_SERVER});

	fetch(constants.API_DOMAIN + constants.API_GET_WATCH_LIST + "?email=" + email, {
		method: 'GET',
		headers: constants.API_JSON_HEADER
	}).then((response) => response.json())
		.then((data) => {
			if (!data || data.length == 0) {
				return false;
			}
			console.error(data);
			var quoteData;
			if(data.watchList.query.count > 1){
				quoteData = data.watchList.query.results.quotes;
			}else{
				quoteData = data.watchList.query.results.quote;
			}
			dispatcher.dispatch({
				type: constants.RECEIVED_WATCH_LIST,
				data: quoteData
			})
		}
	);
}