import { Injectable } from '@angular/core';
import { HttpClient } from '../../node_modules/@angular/common/http';
import { Observable } from '../../node_modules/rxjs';
import { catchError } from '../../node_modules/rxjs/operators';
import { AuthService } from './auth.service';
import { User } from './user';

const ACCOUNT_ENDPOINT = '/api/accounts';

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  constructor(private http: HttpClient, private authService: AuthService) { }

  getAccounts(): Observable<User[]> {
    return this.http.get<User[]>(ACCOUNT_ENDPOINT)
      .pipe(
        catchError(this.authService.handleHttpError.bind(this.authService)),
    );
  }

}
