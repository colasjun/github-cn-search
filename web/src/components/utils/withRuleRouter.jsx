// import withRule from './withRule';
import { flowRight } from 'lodash';
import { withRouter } from 'react-router';

export default function withRuleRouter(WrappedComponent) {
    return flowRight(withRouter)(WrappedComponent);
}