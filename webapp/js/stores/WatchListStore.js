/**
 * Created by yen-chieh on 5/25/16.
 */
import {EventEmitter} from 'events';
import dispatcher from '../dispatcher/appDispatcher.js'
import constant from '../constants/mainConstants.js'

class WatchListStore extends EventEmitter {
	constructor() {
		super();
		this.watchList = localStorage.getItem(constant.LOCAL_WATCH_LIST) ? JSON.parse(localStorage.getItem(constant.LOCAL_WATCH_LIST)) : [];
	}

	updateWatchList(data) {
		this.watchList = data;

		localStorage.setItem(constant.LOCAL_WATCH_LIST, JSON.stringify(this.watchList));
		this.emit("RECEIVED_WATCH_LIST");
	}

	getWatchList(){
		return this.watchList;
	}

	handleAction(action) {
		switch(action.type) {
			case constant.RECEIVED_WATCH_LIST:
				this.updateWatchList(action.data);
				break;
		}
	}
}

const watchListStore = new WatchListStore();
dispatcher.register(watchListStore.handleAction.bind(watchListStore));

export default watchListStore;