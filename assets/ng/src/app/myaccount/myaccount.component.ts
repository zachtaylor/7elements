import { Component, OnInit } from '@angular/core'
import { MyAccount } from '../api'
import { ConnService } from '../conn.service'
import { WebsocketService } from '../websocket.service'
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

  changeEmail = new FormGroup({
    email: new FormControl(''),
  })

  changePassword = new FormGroup({
    password1: new FormControl(''),
    password2: new FormControl(''),
  })

  objectKeys(obj : any) {
    if (!obj) return []
    return Object.keys(obj)
  }

  constructor(public conn : ConnService, public ws : WebsocketService, private router : Router) {
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
    this.ws.send('/logout', {})
  }

  onSubmit() {
    if (this.signup) {
      this.ws.send('/signup', {
        username:this.form.get('username').value,
        email:this.form.get('email').value,
        password1:this.form.get('password1').value,
        password2:this.form.get('password2').value
      })
    } else {
      this.ws.send('/login', {
        username:this.form.get('username').value,
        password:this.form.get('password1').value
      })
    }
  }

  onSubmitChangeEmail() {
    let inputEmail = this.changeEmail.get('email')
    let newEmail = inputEmail.value
    inputEmail.setValue('')
    this.ws.send('/email', {
      email:newEmail
    })
  }
  onSubmitChangePassword() {
    let inputPassword1 = this.changePassword.get('password1')
    let inputPassword2 = this.changePassword.get('password2')
    let newPassword1 = inputPassword1.value
    let newPassword2 = inputPassword1.value
    inputPassword1.setValue('')
    inputPassword2.setValue('')
    this.ws.send('/password', {
      password1:newPassword1,
      password2:newPassword2,
    })
  }
}
