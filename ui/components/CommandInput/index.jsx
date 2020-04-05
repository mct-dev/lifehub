import Input from '../Input'
import { handleCommand } from '../../commands/parse'

const CommandInput = () => {
  const handleChange = (e) => {
    handleCommand(e.target.value)
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
