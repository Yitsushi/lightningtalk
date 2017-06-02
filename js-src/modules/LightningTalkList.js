class LightningTalkList extends React.Component {
  render() {
    var talks = this.props.talks.map(function(talk) {
      return React.createElement(LightningTalkLine, talk)
    });

    return React.createElement("div", {className: "talk-list"}, talks);
  }
}
