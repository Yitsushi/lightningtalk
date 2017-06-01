class LightningTalk extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return React.createElement(
      "div",
      null,
      React.createElement(AddLightningTalkForm)
    );
  }
}

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
      { onSubmit: this.handleSubmit, method: 'post' },
      React.createElement(
        "input",
        {
          value: this.state.title,
          onChange: this.handleTitleChange
        }
      ),
      React.createElement(
        "textarea",
        {
          value: this.state.description,
          onChange: this.handleDescriptionChange
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
    console.log(this.state.title);
    console.log(this.state.description);
  }
}

ReactDOM.render(React.createElement(AddLightningTalkForm, {}), document.getElementById("rootnode"));

