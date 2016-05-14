/**
 * Created by yen-chieh on 5/13/16.
 */

import dispatcher from '../dispatcher/appDispatcher.js'

export function createList(text) {
		dispatcher.dispatch({
			type: 'CREATE_LIST',
			item: text
		});
}