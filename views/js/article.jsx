export default class Article extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            liked: ""
        }
        this.like = this.like.bind(this)
    }
    like() {
        
    }
    render() {
        return (
            <div className="col-xs-4">
                <div className="panel panel-default">
                    <div className="panel-heading">#{this.props.article.id} <span className="pull-right">{this.state.liked}</span></div>
                    <div className="panel-body">
                        {this.props.article.article}
                    </div>
                    <div className="panel-footer">
                        {this.props.article.likes} Likes &nbsp;
                        <a onClick={this.like} className="btn btn-default">
                            <span className="glyphicon glyphicon-thumbps-up"></span>
                        </a>
                    </div>
                </div>
            </div>
        )
    }
}