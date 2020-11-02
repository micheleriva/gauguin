import * as React from 'react';
import csx from 'classnames';

const cardClasses = csx([
  'rounded', 'shadow-lg',
  'bg-gray-100',
  'w-100',
  'p-6'
]);


const Card: React.FC<{children: JSX.Element | string}> = ({ children }) => {
  return (
    <div className={cardClasses}>
      { children }
    </div>
  )
};

export default Card;