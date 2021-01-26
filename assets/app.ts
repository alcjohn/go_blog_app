import register from 'preact-custom-element';
import ButtonDelete from './components/ButtonDelete';
import Turbolinks from 'turbolinks';
import Form from './components/Form';
import './css/app.scss';
Turbolinks.start();
register(Form, 'x-form', ['action', 'method'], { shadow: false });
register(ButtonDelete, 'x-btn-delete');
console.log('Tugudu');
