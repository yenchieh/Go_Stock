/**
 * Created by yen-chieh on 5/18/16.
 */
var React = require('react');
var update = require('react-addons-update');

var Dialog = React.createClass({
	getInitialState: function(){
		return {
			headerText: "Default Header",
			bodyText: "Default Body",
			additionalBody: "",
			confirmButton: ""
		}
	},

	componentDidMount: function(){
		var self = this;
		this.$dialog = $('#popupDialog');
		this.$dialog.on('hidden.bs.modal', function(){
			self.dialogClosed();
		});

		this.setState({
			headerText: this.props.headerText,
			bodyText: this.props.bodyText,
			additionalBody: this.props.additionalBody,
			confirmButton: this.props.confirmButton
		});

		this.$dialog.modal('show');
	},

	dialogClosed: function(){
		if(this.props.closedCallback){
			this.props.closedCallback()
		}

	},

	confirmClicked: function(){

		if(this.props.confirmCallback){
			this.props.confirmCallback($('div.modal-body', this.$dialog));
		}
	},

	render: function () {
		let renderButton = () => {
			if(this.state.confirmButton){
				return (
					<button type="button" className="btn btn-success" onClick={this.confirmClicked} data-dismiss="modal">{this.state.confirmButton}</button>
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
								{this.state.additionalBody}
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

export default Dialog;