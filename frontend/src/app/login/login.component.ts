import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EMPTY } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { AuthService } from '../auth.service';
import { LoginRequest } from '../login-request';
import { TokenResponse } from '../token-response';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  request: LoginRequest = new LoginRequest();
  submitted: boolean;

  constructor(private authService: AuthService, private router: Router) { }

  ngOnInit() {
  }

  onSubmit() {
    this.submitted = true;

    this.authService.requestToken(this.request).pipe(catchError(err => {
      // TODO: display error somehow
      console.log(err);

      return EMPTY;
    })).subscribe(((response: TokenResponse) => {
      this.authService.login(response.token);
      this.router.navigate(['/']);
    }));
  }

}
