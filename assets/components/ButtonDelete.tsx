import { FunctionComponent, h } from 'preact';
import { useCallback } from 'preact/hooks';
interface IProps {
  url: string;
}
const ButtonDelete: FunctionComponent<IProps> = ({ url }) => {
  const _confirm = useCallback((e: Event) => {
    if (confirm('Are you sure ?')) {
      return;
    }
    e.preventDefault();
  }, []);
  return (
    <form action={url} onSubmit={_confirm} method="POST">
      <input type="hidden" name="_method" value="DELETE" />
      <input type="submit" value="Delete" />
    </form>
  );
};
export default ButtonDelete;
