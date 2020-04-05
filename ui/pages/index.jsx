import Link from 'next/link'
import Modal from '../components/Modal'
import CommandInput from '../components/CommandInput'

export default function IndexPage () {
  return (
    <div>
      <nav>
        <Link href="/dashboard"><a title="Dashboard">Dashboard</a></Link>
      </nav>
      <h1>Home page</h1>
      <Modal>
        <CommandInput />
      </Modal>
    </div>
  )
}
