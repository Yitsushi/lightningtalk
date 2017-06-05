class TalkToolbar extends React.Component {
  render() {
    return React.createElement("div", {className: "toolbar"},
      React.createElement("div", {className: "label"}, "Move"),
      React.createElement("div", {className: "button"}, "Top"),
      React.createElement("div", {className: "button"}, "Up"),
      React.createElement("div", {className: "button"}, "Down"),
      React.createElement("div", {className: "button"}, "Bottom")
    );
  }
}

