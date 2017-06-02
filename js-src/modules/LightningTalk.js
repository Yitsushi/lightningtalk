class LightningTalk extends React.Component {
  constructor(props) {
    super(props);
    this.state = { talks: [] };
  }

  componentDidMount() {
    this.updateTalks();
  }

  updateTalks() {
    ServerRequest("get", "/talk", null, function(data) {
      this.setState({talks: data});
    }.bind(this));
  }

  render() {
    return React.createElement(
      "div",
      null,
      React.createElement(AddLightningTalkForm, {updateTalks: this.updateTalks.bind(this)}),
      React.createElement(LightningTalkList, {talks: this.state.talks})
    );
  }
}
