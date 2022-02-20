<?php
namespace App\Controller;
use Symfony\Component\HttpFoundation\Response;

class ContactController{
    function contact(){
        return new Response("test");
    }
}