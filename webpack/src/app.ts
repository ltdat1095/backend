import page from 'page';

import Login from './Login/index';
import SignUp from './SignUp/index';

page('/login', () => { Login });
page('/signup', () => { SignUp });
