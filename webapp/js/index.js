import React from 'react';
import ReactDOM from 'react-dom';
import "!style!css!sass!../scss/index.scss";
import {Router, Route, IndexRoute, Link, browserHistory} from 'react-router';

import Footer from './components/footer.js';
import SearchPage from './components/search.js';
import UserDialog from './components/SignUserDialog.js';
import UserStore from './stores/UserStore.js';
import WatchList from './components/watchList.js';
import * as QuoteAction from './actions/QuoteAction.js';
import SearchStockStore from './stores/SearchStockStore.js';


var Index = React.createClass({
	getInitialState: function(){
		return {
			user: {},
			watchList: [],
			showSignUserDialog: false
		}
	},

	componentDidMount: function(){
		UserStore.on('userInfoUpdated', this.updateUserState);
	},

	updateUserState: function(){
		this.setState({
			user: UserStore.getUser()
		})
	},

	addListButtonClicked: function(){
		console.error("clicked");
		console.error(SearchStockStore.getCheckedQuote());
		if(this.state.user == {} || !this.state.user.isLoggedIn){
			this.setState({showSignUserDialog: true})
		}else{
			QuoteAction.addCheckedQuoteToWatchList(this.state.user, SearchStockStore.getCheckedQuote());
		}
	},

	userDialogClosed: function(){
		this.setState({showSignUserDialog: false});
	},

	render: function(){

		let renderWatchListTab = function(){
			if(this.state.user.isLoggedIn){
				return (
					<li role="presentation" className="tab" onClick={this.clickedOnTab}><Link to="/watchList">Watch List</Link></li>
				)
			}

		}.bind(this);

		let showUserDialog = () => {
			if (this.state.showSignUserDialog) {
				return (
					<UserDialog closedCallback={this.userDialogClosed} />
				)
			}
		};

		let loggedInUser = () => {
			if(this.state.user.name){
				return(
					<div>
						Logged in as {this.state.user.name} | Logout
					</div>
				)
			}else{
				return (
					<button className="btn btn-sml btn-success loginButton" onClick={this.addListButtonClicked}>Login | Sign-up</button>
				)
			}
		};

		return (
			<div>
				{showUserDialog()}
				{loggedInUser()}
				<div id="nav">
					<ul className="nav nav-tabs">
						<li role="presentation" className="tab active" onClick={this.clickedOnTab}><Link to="/">Search</Link></li>
						{renderWatchListTab()}
					</ul>
				</div>

				{React.cloneElement(this.props.children, {addListButtonClicked: this.addListButtonClicked})}
				<Footer/>
			</div>
		)
	}
});


ReactDOM.render(
	<Router history={browserHistory}>
		<Route path="/" component={Index}>
			<IndexRoute component={SearchPage}/>
			<Router path="/watchList" component={WatchList} />
		</Route>
	</Router>,
	document.getElementById('app')
);
