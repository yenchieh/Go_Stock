/**
 * Created by yen-chieh on 5/18/16.
 *
 * This should use Mixin to reuse Dialog.js
 * https://facebook.github.io/react/docs/reusable-components.html
 */

var React = require('react');
var update = require('react-addons-update');
import UserStore from '../stores/UserStore.js';
import * as UserAction from '../actions/UserAction.js';

var SignUserDialog = React.createClass({
	getInitialState: function () {
		return {
			headerText: "Register / Login Account",
			bodyText: "Please Register or Login to store your watch list.",
			confirmButton: "Register | Login",
			isRegister: false
		}
	},

	componentDidMount: function () {
		var self = this;
		this.$dialog = $('#popupDialog');
		this.$email = $('#userEmailInput');
		this.$password = $('#userPasswordInput');

		this.$dialog.on('hidden.bs.modal', function () {
			self.props.closedCallback();
		});

		this.$dialog.modal('show');
	},

	submitClicked: function () {
		var self = this;
		this.$dialog.button('loading');
		UserAction.RegisterOrLogin(this.$email.val(), this.$password.val(), $('#userNameInput').val());

		UserStore.on('userInfoUpdated', function(){
			self.$dialog.button('reset');
			self.$dialog.modal('hide');
		});
	},

	setUserExistInDatabase: function () {
		this.setState({
			headerText: "Login Account",
			bodyText: "Please Login to store your watch list.",
			confirmButton: "Login"
		})
	},

	setUserNotExistInDatabase: function () {
		this.setState({
			headerText: "Register Account",
			bodyText: "Please Register to store your watch list.",
			confirmButton: "Register"
		})
	},

	checkUserExist: function (e) {
		UserStore.on('userEmailChecked', () => {
			if (UserStore.isUserExistInDatabase()) {
				this.setUserExistInDatabase();
			} else {
				this.setUserNotExistInDatabase();
			}
			this.setState({
				isRegister: !UserStore.isUserExistInDatabase()
			})
		});
		UserAction.CheckUserExists(e.target.value);
	},

	render: function () {
		let renderButton = () => {
			if (this.state.confirmButton) {
				return (
					<button type="button" className="btn btn-primary" onClick={this.submitClicked}>{this.state.confirmButton}</button>
				)
			}
		};

		let nameInput = () => {
			if(this.state.isRegister){
				return (
					<input type="text" name="name" className="form-control" id="userNameInput" placeholder="Name"/>
				)
			}
		};

		return (
			<div className="container">

				<div id="popupDialog" className="modal fade" role="dialog">
					<div className="modal-dialog">

						<div className="modal-content">
							<div className="modal-header">
								<button type="button" className="close" data-dismiss="modal">&times;</button>
								<h4 className="modal-title">{this.state.headerText}</h4>
							</div>
							<div className="modal-body">
								<p>{this.state.bodyText}</p>

								<div className="accountDialog">
									<input type="email" name="email" className="form-control" id="userEmailInput"
												 placeholder="Email" onBlur={this.checkUserExist}/><br/>
									<input type="password" name="password" className="form-control" id="userPasswordInput"
												 placeholder="Password"/><br/>
									{nameInput()}

								</div>
							</div>
							<div className="modal-footer">
								{renderButton()}
							</div>
						</div>

					</div>
				</div>
			</div>

		)
	}
});

export default SignUserDialog;