/**
 * Created by yen-chieh on 5/25/16.
 */

import React from 'react';

var WatchListTable = React.createClass({
	getInitialState: function(){
		return{
			quotes: []
		}
	},

	componentWillReceiveProps: function(props){
		this.setState({
			quotes: props.quote
		})
	},

	renderColoredStockChange: function(change, percentChange){
		let className = change < 0 ? "negative" : "positive";
		return(
			<div>
				<span className={className}>{change}</span> / <span className={className}>{percentChange}</span>
			</div>

		)
	},

	render: function(){

		var quoteTable = this.state.quotes.map(function (quote, i) {
			return (
				<tr key={i}>
					<td>
						{quote.name}
					</td>
					<td>
						<div className="symbol">{quote.symbol}</div>
					</td>
					<td>
						{quote.open}
					</td>
					<td>
						{this.renderColoredStockChange(parseFloat(quote.change).toFixed(2), parseFloat(quote.percentChange).toFixed(2))}
					</td>
					<td>
						{parseFloat(quote.daysLow).toFixed(2)} / {parseFloat(quote.daysHigh).toFixed(2)}
					</td>
					<td>
						{quote.volume}
					</td>
				</tr>
			)

		}.bind(this));


		return (
			<div className="quote">
				<table className="table table-hover">
					<thead>
					<tr>
						<th>Name</th>
						<th>Symbol</th>
						<th>Open</th>
						<th>Change / Percent</th>
						<th>DaysLow / DaysHigh</th>
						<th>Volume</th>
					</tr>
					</thead>

					<tbody>
					{quoteTable}

					</tbody>
				</table>

			</div>
		)
	}
});

export default WatchListTable;