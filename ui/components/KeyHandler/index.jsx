import PropTypes from 'prop-types'

const KeyHandler = ({ keyCombo, handle }) => {
  return null
}
KeyHandler.propTypes = {
  keyCombo: PropTypes.arrayOf(PropTypes.string).isRequired,
  handle: PropTypes.func
}

export default KeyHandler
