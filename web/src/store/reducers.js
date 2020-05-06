
import { combineReducers } from 'redux'
import defaultState from './state.js'
function pageTitle (state = defaultState.pageTitle, action) {
  switch (action.type) {
    case 'SET_PAGE_TITLE':
      return action.data
    default:
      return state
  }
}

export default combineReducers({
    pageTitle 
})