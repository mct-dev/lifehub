import PropTypes from 'prop-types'

export default function Input ({ type, onChange, placeholder = '', ...rest }) {
  return (
    <input
      style={inputStyle}
      type={type}
      onChange={onChange}
      placeholder={placeholder}
      {...rest}
    />
  )
}
Input.propTypes = {
  type: PropTypes.oneOf(['text', 'number']).isRequired,
  onChange: PropTypes.func,
  placeholder: PropTypes.string
}

const inputStyle = {
  padding: '10px 5px',
  border: '1px solid #e8e8e8',
  borderRadius: '3px'
}
