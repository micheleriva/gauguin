import * as React from 'react';
import { useLocation, Link } from 'react-router-dom';
import csx from 'classnames';

const navClasses = csx([
  'pb-5', 'mb-5',
  'border-b-2', 'border-gray-300'
]);

const navbarClasses = csx([
  'flex'
]);

const navbarItemClasses = csx([
  'mr-3'
]);

const navbarLinkClasses = (currentPath: string, path: string ): string => csx([
  'inline-block', 'rounded', 'py-1', 'px-3',
  { 'bg-purple-500 text-white': currentPath === path },
  { 'text-purple-500': currentPath !== path }
]);

const Navbar = () => {
  const { pathname } = useLocation();

  return (
    <nav className={navClasses}>
      <ul className={navbarClasses}>
        <li className={navbarItemClasses}>
          <Link to='/' className={navbarLinkClasses(pathname, '/')}>
            Home
          </Link>
        </li>
        <li className={navbarItemClasses}>
          <Link to='/generate' className={navbarLinkClasses(pathname, '/generate')}>
            Generate
          </Link>
        </li>
        <li className={navbarItemClasses}>
          <Link to='/docs' className={navbarLinkClasses(pathname, '/docs')}>
            Docs
          </Link>
        </li>
        <li className={navbarItemClasses}>
          <a className={navbarLinkClasses('', '-')} href='https://github.com/micheleriva/gauguin' target='__blank'>
            GitHub
          </a>
        </li>
      </ul>
    </nav>
  )
};

export default Navbar;