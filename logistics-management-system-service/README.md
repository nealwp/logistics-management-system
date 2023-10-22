# Logistics Management System Service


* No allowance can exist without an item
* No location can exist without an item
* No stock requisition can exist without an item
* No stock requisition can be created without an allowance

## Requirements

* ~~Add an Item~~
  * ~~id number~~
  * ~~name~~
  * ~~price~~
  * should not create if item already exists

* Remove an Item
  * fails if item does not exist
  * fails if inventory exists
  * fails if orders exists 
  * fails if requisitions exist

* Establish allowance
  * ~~fails if item does not exist~~
  * ~~provide quantity~~

* Create a stock requisition
  * unique identifier
  * item record must exist
  * item must have allowance

* Initiate stock-fill buy
  * creates stock requisitions for all items under allowance level

* Receive stock requisition
  * sends receipt status to supplier
  * creates pending stow quantity

* Assign location
  * fails if item does not exist

* Stow item
  * requires inventory location
  * requires pending stow
  * updates on-hand quantity
  * update pending stow quantity

* Search item
  * show id, name, and price
  * show allowance level, quantity on order, pending stow, and onhand
  * show location

* Search requisition
  * show req id, item ordered, date ordered
  * show quantity ordered, total cost
  * show req status, status date
  * show recieved date or cancellation date

* Event Log
  * append-only transaction ledger
  * every action on item creates event
  * tracks date and any pertinent info

* Initiate inventory
  * places stock items under inventory

* Post inventory
  * updates inventory on-hand
  * creates loss/gain if new quantity differs from current quantity


