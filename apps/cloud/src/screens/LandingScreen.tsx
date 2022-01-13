import type { FC } from 'react'
import { UserIcon, DatabaseIcon } from '@heroicons/react/outline'
import { ApiBox } from '../components/ApiBox'
import { FeedbackButton } from '../components/FeedbackButton'

export const LandingScreen: FC = () => {
  return (
    <section className="p-14 text-white">
      <h1 className="text-5xl font-bold mb-4">Welcome to M3O Cloud</h1>
      <p className="text-xl max-w-xl text-gray-400">
        Manage your content on M3O with M3O Cloud!
      </p>
      <div className="mt-10 grid grid-cols-4 gap-4">
        <ApiBox
          title="Users"
          subTitle="Manage your users"
          icon={<UserIcon className="w-6" />}
          linkTo="/users"
        />
        <ApiBox
          title="DB"
          subTitle="Manage your data"
          icon={<DatabaseIcon className="w-6" />}
          linkTo="/database"
        />
      </div>
      <p className="mt-8 max-w-2xl border-t pt-6 text-gray-400 border-gray-700">
        We're always looking for ways to improve our customer experience. Please
        provide feedback or any feature requests you have for this product.
      </p>
      <p>
        <FeedbackButton />
      </p>
    </section>
  )
}
