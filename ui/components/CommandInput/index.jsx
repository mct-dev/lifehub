import Input from '../Input'
import { handleCommand } from '../../commands/handler'

const CommandInput = () => {
  const handleChange = (e) => {
    handleCommand(e.target.value.trim())
  }

  return (
    <div style={styles}>
      <Input type="text" onChange={handleChange} placeholder='Enter command...' />
    </div>
  )
}

const styles = {
  width: '100%',
  display: 'flex',
  justifyContent: 'center',
  marginTop: '40px'
}

export default CommandInput
