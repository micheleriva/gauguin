import * as React from 'react';
import { Link } from 'react-router-dom';
import csx from 'classnames';
import Card from '../components/card';
import { Button } from '../components/button';

const cardTitleClasses = csx([
  'font-bold', 'text-xl'
]);

const cardDescriptionClasses = csx([
  'text-gray-600', 'mt-2', 'mb-4'
]);

const ctaClasses = [
  'bg-blue-500', 'hover:bg-blue-600', 'mr-5'
];

const Home = () => {
  return (
    <div className='grid grid-cols-3 p-10 gap-10'>
      <Card>
        <>
          <h2 className={cardTitleClasses}> Getting Started </h2>
          <p className={cardDescriptionClasses}> Read the documentation about <b>Gauguin</b> </p>
          <Link to='/docs'>
            <Button classes={ctaClasses}>
              <span className={'text-white'}>
                Let's start!
              </span>
            </Button>
          </Link>
        </>
      </Card>
      <Card>
        <>
          <h2 className={cardTitleClasses}> Generate Images </h2>
          <p className={cardDescriptionClasses}> Generate Images based on your configuration file. </p>
          <Link to='/generate'>
            <Button classes={ctaClasses}>
              <span className={'text-white'}>
                Let's create!
              </span>
            </Button>
          </Link>
        </>
      </Card>
      <Card>
        <>
          <h2 className={cardTitleClasses}> Contribute </h2>
          <p className={cardDescriptionClasses}> <b>Gauguin</b> it's a free open source software. </p>
          <a href='https://www.github.com/micheleriva/gauguin' target='__blank'>
            <Button classes={ctaClasses}>
              <span className={'text-white'}>
                Contribute on GitHub
              </span>
            </Button>
          </a>
        </>
      </Card>
    </div>
  )
};

export default Home;