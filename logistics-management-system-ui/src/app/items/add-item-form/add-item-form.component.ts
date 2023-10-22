import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-add-item-form',
  templateUrl: './add-item-form.component.html',
  styleUrls: ['./add-item-form.component.scss']
})
export class AddItemFormComponent {

  item = {
      id: '',
      name: '',
      unitPrice: 0
  }

  constructor(private router: Router){};

  public onSubmit() {
      console.log("submitted")
      this.router.navigate(['/']);
  }

  public onCancel() {
      this.router.navigate(['/']);
  }
}
