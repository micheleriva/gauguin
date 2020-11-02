import * as React from 'react';
import qs from 'querystring';
import csx from 'classnames';
import { DownloadButton, Button } from '../components/button';

const contClasses = csx([
  'grid', 'gap-10'
]);

const labelClasses = csx([
  'block', 'text-gray-700', 'text-sm', 'font-bold', 'mb-2',
  'capitalize'
]);

const inputClasses = csx([
  'shadow', 'appearance-none', 'border', 'rounded',
  'w-full', 'py-2', 'px-3', 'text-gray-700',
  'leading-tight', 'focus:outline-none', 'focus:shadow-outline'
]);

const imgContainerClasses = csx([]);

const ctaContClasses = csx([
  'pt-5'
]);

const selectContClasses = csx([
  'inline-block', 'relative', 'w-64', 'mb-5'
]);

const selectClasses = csx([
  'block', 'appearance-none', 'w-full', 
  'bg-white', 'border', 'border-gray-400', 
  'hover:border-gray-500', 'px-4', 'py-2', 
  'pr-8', 'rounded', 'shadow', 'leading-tight',
  'focus:outline-none', 'focus:shadow-outline'
]);

const Generate = () => {
  // @ts-ignore
  const config = window.__gauguin_config;
  const routes = config.routes;
  const [route, setRoute] = React.useState(routes[1]);

  const [paramStates, setParamStates] = React.useState({})
  const handleSetParam = (event: React.ChangeEvent<HTMLInputElement>, param) =>
    setParamStates({ ...paramStates, [param]: event?.target?.value ?? '' })

  const queryString = qs.stringify(paramStates);
  const imageUrl = `http://localhost:8080${route.path}?${queryString}`;

  React.useEffect(() => {
    setParamStates({});
  }, route.path);

  return (
    <div className={contClasses} style={{ gridTemplateColumns: '450px 1fr' }}>
      <div>
        <div className={selectContClasses}>
          <select className={selectClasses} onChange={console.log}>
            {routes.map((route) => <option key={route.name} value={route}> {route.name} </option>)}
          </select>
          <div className='pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700'>
            <svg className='fill-current h-4 w-4' xmlns='http://www.w3.org/2000/svg' viewBox='0 0 20 20'><path d='M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z'/></svg>
          </div>
        </div>
        <div>
          {
            route.params.map(param => (
              <div className='mb-3' key={param}>
                <label className={labelClasses} htmlFor={param}>
                  { param }
                </label>
                <input
                  className={inputClasses}
                  id={param}
                  type='text'
                  placeholder={ param }
                  onChange={(event) => handleSetParam(event, param)}
                >
                </input>
              </div>
            ))
          }
        </div>
        <div className={ctaContClasses}>
          <a href={imageUrl} target='__blank'>
            <Button classes={['bg-blue-500', 'hover:bg-blue-600', 'mr-5']}>
              <span className={'text-white'}>
                View Full Size
              </span>
            </Button>
          </a>
          <a href={imageUrl} download={route.name} target='__blank'>
            <DownloadButton />
          </a>
        </div>
      </div>
      <div className={imgContainerClasses}>
        <img
          src={imageUrl}
          style={{ width: '100%' }}
          alt=''
        />
      </div>
    </div>
  )
}

export default Generate;