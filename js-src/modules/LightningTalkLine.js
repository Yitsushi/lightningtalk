class LightningTalkLine extends React.Component {
  claim() {
    console.log(this.props.Id);
  }

  render() {
    var owner = "open";
    var ownerClass = "owner";

    if (this.props.Presenter === null) {
      ownerClass += " open";
    }

    return React.createElement(
      "div", {className: "talk-line"},
      React.createElement("div", {className: "title"}, this.props.Title),
      React.createElement("div", {className: "description"}, this.props.Description),
      React.createElement("div", {onClick: this.claim.bind(this), className: ownerClass}, owner),
      React.createElement(TalkToolbar, this.props)
    );
  }
}
