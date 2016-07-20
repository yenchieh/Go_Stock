/*jshint esversion: 6 */
/**
 * Created by yen-chieh on 5/12/16.
 */

import {EventEmitter} from 'events';
import dispatcher from '../dispatcher/appDispatcher.js'
import constants from '../constants/mainConstants.js';

class listStore extends EventEmitter {
	constructor(){
		super();
		this.list = [
			"1", "2", "3"
		]
	}

	createList(text) {
		this.list.push(text);
		this.emit("change");
	}

	getAll(){
		return this.list;
	}

	handleAction(action) {
		switch(action.type){
			case constants.CREATE_LIST:
				this.createList(action.item);
				break;
		}
	}
}

const ListStore = new listStore();
dispatcher.register(ListStore.handleAction.bind(ListStore));

export default ListStore;
