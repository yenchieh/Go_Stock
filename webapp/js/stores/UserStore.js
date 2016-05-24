/**
 * Created by yen-chieh on 5/18/16.
 */

import {EventEmitter} from 'events';
import dispatcher from '../dispatcher/appDispatcher.js'
import constant from '../constants/mainConstants.js'
import SearchStockStore from './SearchStockStore.js';

class UserStore extends EventEmitter {
	constructor() {
		super();
		this.user = {};
		this.watchList = [];
		this.emailExistInDatabase = false;
	}

	isUserExistInDatabase(){
		return this.emailExistInDatabase;
	}

	checkedUserEmail(result) {
		this.emailExistInDatabase = result.exist == true;
		this.emit("userEmailChecked");
	}

	updateUserInfo(result){
		this.user = {
			email: result.email,
			name: result.name,
			isLoggedIn: true
		};

		this.emit("userInfoUpdated");
	}

	getUser(){
		return this.user;
	}

	watchListUpdated(data){
		this.watchList = data.watchList;
		SearchStockStore.removeQuoteByCheckedList();

		this.emit("watchListUpdated");
	}

	handleAction(action) {
		switch(action.type){
			case constant.USER_REGISTER_LOGIN_SUCCESS:
				this.updateUserInfo(action.data);
				break;
			case constant.CHECKED_USER_EXISTS:
				this.checkedUserEmail(action.data);
				break;
			case constant.RECEIVED_QUOTE_DATA_WATCH:
				break;
			case constant.WATCH_LIST_UPDATED:
				this.watchListUpdated(action.data);
				break;
		}
	}
}

const userStore = new UserStore();
dispatcher.register(userStore.handleAction.bind(userStore));

export default userStore;