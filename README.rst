===============
Mongodb Handson
===============

Usage
=====

.. code-block:: bash

   # start mongodb
   $ docker-compose up


   # if you change configrations
   $ docker-compose down && docker-compose build && docker-compose up


   # connect to mongo
   $ mongo localhost:27018

   # connect to initialized database
   $ mongo localhost:27018/tests -u gopher -psecret
