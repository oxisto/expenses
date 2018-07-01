import { Component, OnInit } from '@angular/core';
import { LoginRequest } from '../login-request';
import { AuthService } from '../auth.service';
import { catchError } from 'rxjs/operators';
import { empty } from 'rxjs';
import { TokenResponse } from '../token-response';
import { Router } from '@angular/router';

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
      return empty();
    })).subscribe(((response: TokenResponse) => {
      this.authService.login(response.token);
      this.router.navigate(['/']);
    }));
  }

}
