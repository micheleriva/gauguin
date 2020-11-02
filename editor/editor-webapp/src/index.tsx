import React from 'react';
import ReactDOM from 'react-dom';
import { Switch, Route, HashRouter } from 'react-router-dom';
import csx from 'classnames';
import Home from './pages/Home';
import Generate from './pages/Generate';
import Navbar from './components/navbar'
import reportWebVitals from './reportWebVitals';
import './assets/tailwind.output.css'

const containerClasses = csx([
  'flex', 'items-center', 'justify-center',
  'bg-gray-300',
  'min-w-full', 'h-screen',
]);

const cardClasses = csx([
  'rounded', 'shadow-lg',
  'bg-white',
  'w-9/12',
  'p-10'
]);

const headerClasses = csx([
  'text-4xl',
  'font-bold',
  'w-full',
  'text-center',
  'mb-10'
]);

ReactDOM.render(
  <div className={containerClasses}>
    <div className={cardClasses}>
      <h1 className={headerClasses}> GAUGUIN </h1>
      <React.StrictMode>
        <HashRouter>
          <Navbar />
          <Switch>
            <Route exact path="/">
              <Home />
            </Route>
            <Route exact path="/generate">
              <Generate />
            </Route>
          </Switch>
        </HashRouter>
      </React.StrictMode>
    </div>
  </div>,
  document.getElementById('root')
);

reportWebVitals();
