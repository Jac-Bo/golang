import {LoggedIn} from "./login"
import {Home} from './home'

export default class App extends React.Component {
  render() {
    if (this.loggedIn) {
        return (<LoggedIn />);
    } else {
        return (<Home />);
    }
}
}
ReactDOM.render(<App />, document.getElementById('app'));
