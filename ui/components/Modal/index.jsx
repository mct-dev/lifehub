import PropTypes from 'prop-types'

export default function Modal ({ children }) {
  return (
    <div style={modalStyles}>
      {children}
    </div>
  )
}
Modal.propTypes = {
  children: PropTypes.node.isRequired
}

const modalStyles = {
  position: 'fixed',
  top: '0',
  left: '0',
  width: '100vw',
  height: '100vh',
  padding: '20px',
  background: '#fff',
  zIndex: '10000'
}
