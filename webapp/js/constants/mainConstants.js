/**
 * Created by yen-chieh on 5/11/16.
 */

module.exports = {
	API_DOMAIN: "http://localhost:8080/",

	API_REMOVE_LIST: "removeFromList",
	API_USER_STOCK_LIST: "getUserStockList",
	API_SEARCH_STOCK: "getQuoteBySymbol",
	API_ADD_STOCK: "addToList",
	API_CHECK_USER: "checkUser",

	API_JSON_HEADER: {
		'Accept': 'application/json',
		'Content-Type': 'application/json'
	},

	FETCH_DATA_FROM_SERVER: 'FETCH_DATA',
	RECEIVED_QUOTE_DATA: 'RECEIVED_QUOTE_DATA',
	UPDATE_CHECKED_LIST: 'UPDATE_CHECKED_LIST',
	REMOVE_QUOTE: 'REMOVE_QUOTE',
	GET_QUOTE: 'GET_QUOTE',
	CREATE_LIST: 'CREATE_LIST'
};