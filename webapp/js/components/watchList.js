/**
 * Created by yen-chieh on 5/12/16.
 */

import React from 'react';
import _ from 'underscore';
import dispatcher from '../dispatcher/appDispatcher.js';
import constant from '../constants/mainConstants.js';
import WatchListTable from './watchListQuote.js';

import WatchListStore from '../stores/WatchListStore.js';
import * as WatchListAction from '../actions/WatchListAction.js';
import UserStore from '../stores/UserStore.js';

var WatchListPage = React.createClass({
	getInitialState: function () {
		return {
			watchList: []
		}
	},

	componentWillMount: function () {
		var self = this;
		WatchListStore.on('RECEIVED_WATCH_LIST', function(){
			self.updateWatchList();
		});
	},

	componentDidMount: function () {
		WatchListAction.getWatchList(UserStore.getUser().email);
		//this.updateWatchList();
	},

	updateWatchList: function(){
		this.setState({
			watchList: WatchListStore.getWatchList()
		});
	},

	render: function () {
		return (
			<div>
				<h1>Watch List</h1>
				<WatchListTable quote={this.state.watchList} />
			</div>
		)
	}
});

export default WatchListPage;