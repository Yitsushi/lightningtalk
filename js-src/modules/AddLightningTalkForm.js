class AddLightningTalkForm extends React.Component {
  constructor(props) {
    super(props);
    this.handleTitleChange = this.handleTitleChange.bind(this);
    this.handleDescriptionChange = this.handleDescriptionChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.state = { title: '', description: '' };
  }

  render() {
    return React.createElement(
      "form",
      { onSubmit: this.handleSubmit, method: 'post', className: "add-new-form" },
      React.createElement(
        "input",
        {
          value: this.state.title,
          onChange: this.handleTitleChange,
          placeholder: "Title"
        }
      ),
      React.createElement(
        "textarea",
        {
          value: this.state.description,
          onChange: this.handleDescriptionChange,
          placeholder: "Description"
        }
      ),
      React.createElement("button", null, "Send")
    );
  }

  handleTitleChange(e) {
    this.setState({ title: e.target.value });
  }

  handleDescriptionChange(e) {
    this.setState({ description: e.target.value });
  }

  handleSubmit(e) {
    e.preventDefault();
    var data = {
      Title: this.state.title,
      Description: this.state.description
    };
    return ServerRequest("post", "/talk", data, function(d) {
      this.setState({title: "", description: ""});
      return this.props.updateTalks();
    }.bind(this))
  }
}
