import { Component, OnInit } from '@angular/core'
import { MyAccount } from '../api'
import { ConnService } from '../conn.service'
import { Subscription } from 'rxjs'
import { FormGroup, FormControl } from '@angular/forms'
import { Router } from '@angular/router'

@Component({
  selector: 'app-myaccount',
  templateUrl: './myaccount.component.html',
  styleUrls: ['./myaccount.component.css']
})
export class MyAccountComponent implements OnInit {
  public signup : boolean // signup mode when not logged in

  myaccount : MyAccount
  private $myaccount : Subscription

  form = new FormGroup({
    username: new FormControl(''),
    email: new FormControl(''),
    password1: new FormControl(''),
    password2: new FormControl(''),
  })

  objectKeys(obj : any) {
    if (!obj) return []
    return Object.keys(obj)
  }

  constructor(public conn : ConnService, private router : Router) {
  }

  ngOnInit() {
    this.$myaccount = this.conn.myaccount$.subscribe(myaccount => {
      this.myaccount = myaccount
    })
  }

  ngOnDestroy() {
    this.$myaccount.unsubscribe()
  }

  logout() {
    this.conn.sendWS('/logout', {})
  }

  onSubmit() {
    if (this.signup) {
      this.conn.sendWS('/signup', {
        username:this.form.get('username').value,
        email:this.form.get('email').value,
        password1:this.form.get('password1').value,
        password2:this.form.get('password2').value
      })
    } else {
      this.conn.sendWS('/login', {
        username:this.form.get('username').value,
        password:this.form.get('password1').value
      })
    }
  }
}
