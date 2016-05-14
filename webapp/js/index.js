import React from 'react';
import ReactDOM from 'react-dom';
import "!style!css!sass!../scss/index.scss";
import {Router, Route, IndexRoute, Link, browserHistory} from 'react-router';

import Footer from './components/footer.js';
import SearchPage from './components/search.js';


var Index = React.createClass({


	render: function(){

		let renderTab = function(){
			return (
				<li role="presentation" className="tab" onClick={this.clickedOnTab}></li>
			)
		}.bind(this);

		return (
			<div>
				<div id="nav">
					<ul className="nav nav-tabs">
						<li role="presentation" className="tab active" onClick={this.clickedOnTab}><Link to="/">Search</Link></li>
						{renderTab()}
					</ul>
				</div>

				{this.props.children}
				<Footer/>
			</div>
		)
	}
});


ReactDOM.render(
	<Router history={browserHistory}>
		<Route path="/" component={Index}>
			<IndexRoute component={SearchPage}/>
		</Route>
	</Router>,
	document.getElementById('app')
);
