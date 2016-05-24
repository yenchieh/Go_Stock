/**
 * Created by yen-chieh on 5/12/16.
 */

import React from 'react';
import _ from 'underscore';
import dispatcher from '../dispatcher/appDispatcher.js';
import constant from '../constants/mainConstants.js'

var WatchListPage = React.createClass({
	getInitialState: function () {
		return {
			watchList: []
		}
	},

	componentWillMount: function () {

	},

	componentDidMount: function () {

	},

	render: function () {
		return (
			<div>
				<h1>Watch List</h1>
			</div>
		)
	}
});

export default WatchListPage;