import * as React from 'react';
import csx from 'classnames';

const basicButtonClasses = (...args: any[]) => csx([
  'bg-gray-300', 'hover:bg-gray-400',
  'text-gray-800', 'font-bold py-2', 'px-4',
  'rounded', 'inline-flex', 'items-center',
  ...args
]);

export const DownloadButton = () => {
  return (
    <button className={basicButtonClasses()}>
      <svg className="fill-current w-4 h-4 mr-2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
        <path d="M13 8V2H7v6H2l8 8 8-8h-5zM0 18h20v2H0v-2z"/>
      </svg>
      <span>Download</span>
    </button>
  )
}

type ButtonProps = {
  children: JSX.Element,
  classes: string[]
};

export const Button: React.FC<ButtonProps> = ({ children, classes }) => {
  return (
    <button className={basicButtonClasses(classes)}>
      {children}
    </button>
  )
}