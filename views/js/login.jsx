import {Article} from "./article"

export default class LoggedIn extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            articles: []
        }
    }
    render() {
        return (
            <div className="container">
                <div className="col-lg-12">
                    <br /> 
                    <span className="pull-right"><a onClick={this.logout}>Log Out</a></span>
                    <h2>Artigo</h2>
                    <p>Let's Get Some New Article</p>
                    <div className="row">
                        {this.state.articles.map(function(article, i){
                            return (<Article key={i} article={article} />);
                        })}
                    </div>
                </div>
            </div>
        )             
    }    
}