/**
 * Created by yen-chieh on 5/13/16.
 */

import dispatcher from '../dispatcher/appDispatcher.js';
import constants from '../constants/mainConstants.js';
import 'whatwg-fetch';

export function GetUserData(email){
	dispatcher.dispatch({type: constants.FETCH_DATA_FROM_SERVER});

	fetch(constants.API_DOMAIN + constants.API_RETRIEVE_USER + email, {
		method: 'GET',
		headers: constants.API_JSON_HEADER
	}).then((response) => response.json())
		.then((data) => {
			if (!data || data.length == 0) {
				dispatcher.dispatch({
					type: constants.RECEIVED_USER_DATA,
					data: "No data"
				});
				return false;
			}

			dispatcher.dispatch({
				type: constants.RECEIVED_USER_DATA,
				data: data
			})
		}
	);
}
export function CheckUserExists(email){
	dispatcher.dispatch({type: constants.FETCH_DATA_FROM_SERVER});

	fetch(constants.API_DOMAIN + constants.API_CHECK_EMAIL + email, {
		method: 'GET',
		headers: constants.API_JSON_HEADER
	}).then((response) => response.json())
		.then((data) => {
			dispatcher.dispatch({
				type: constants.CHECKED_USER_EXISTS,
				data: data
			})

		}
	)
}

export function RegisterOrLogin(email, password, name){
	dispatcher.dispatch({type: constants.FETCH_DATA_FROM_SERVER});

	var data = {
		email: email,
		password: password,
		name: name
	};

	fetch(constants.API_DOMAIN + constants.API_REGISTER_LOGIN_USER , {
		method: 'POST',
		headers: constants.API_JSON_HEADER,
		body: JSON.stringify(data)
	}).then((response) => response.json())
		.then((result) => {
			if(result.success){
				dispatcher.dispatch({
					type: constants.USER_REGISTER_LOGIN_SUCCESS,
					data: result.data
				})
			}else{
				dispatcher.dispatch({
					type: constants.ERROR_USER_REGISTER_LOGIN,
					data: result.data
				})
			}


		}
	)
}