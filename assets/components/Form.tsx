import { useCallback } from 'preact/hooks';
import { FunctionComponent, h } from 'preact';

interface IProps {
  action: string;
  method: string;
}
const Form: FunctionComponent<IProps> = ({ action, method, children }) => {
  console.log({
    action,
    method,
    children,
  });
  const _handleSubmit = useCallback(async (e: Event) => {
    e.preventDefault();
    console.log('coucou');
    const form = e.target as HTMLFormElement;
    const formData = { ...Object.fromEntries(new FormData(form)) };
    // const data = await fetchApi(action, {
    //   method,
    //   body: formData,
    // });
    console.log(formData);
  }, []);
  console.log(children);
  return <form onSubmit={_handleSubmit}>{children}</form>;
};
export default Form;
