Terraform es una herramienta de infraestructura como código (IaC) que te permite definir y gestionar tu infraestructura en la nube de manera declarativa. Con Terraform, puedes describir tu infraestructura deseada en archivos de configuración y luego aplicar esos archivos para crear y gestionar recursos en proveedores de nube como AWS.

### Paso a Paso para Usar Terraform con AWS para una API REST y PostgreSQL

#### 1. **Instalar Terraform**: Asegúrate de tener Terraform instalado en tu máquina. Puedes descargarlo desde el [sitio oficial de Terraform](https://www.terraform.io/downloads.html).

Terraform ofrece una variedad de comandos para gestionar la infraestructura. Aquí te dejo una lista de los comandos más comunes y sus usos:

1. **`terraform init`**: Inicializa un directorio de trabajo que contiene archivos de configuración de Terraform. Este comando descarga los proveedores necesarios y prepara el entorno para ejecutar otros comandos de Terraform.

2. **`terraform plan`**: Crea un plan de ejecución, mostrando los cambios que Terraform realizará en la infraestructura. Es útil para revisar los cambios antes de aplicarlos.

3. **`terraform apply`**: Aplica los cambios necesarios para alcanzar el estado deseado de la infraestructura, tal como se define en los archivos de configuración. Este comando ejecuta el plan de cambios.

4. **`terraform destroy`**: Elimina todos los recursos gestionados por Terraform en el directorio de trabajo actual. Es útil para limpiar la infraestructura cuando ya no se necesita.

5. **`terraform validate`**: Verifica que los archivos de configuración de Terraform sean sintácticamente válidos y coherentes.

6. **`terraform fmt`**: Formatea los archivos de configuración de Terraform para que sigan un estilo de codificación estándar.

7. **`terraform show`**: Muestra el estado o el plan de ejecución de Terraform en un formato legible.

8. **`terraform output`**: Muestra los valores de salida definidos en la configuración de Terraform.

9. **`terraform refresh`**: Actualiza el estado de Terraform para reflejar el estado real de la infraestructura.

10. **`terraform import`**: Importa recursos existentes en la infraestructura a Terraform para que puedan ser gestionados.

11. **`terraform taint`**: Marca un recurso gestionado por Terraform para que sea destruido y recreado en la próxima ejecución de `terraform apply`.

12. **`terraform untaint`**: Desmarca un recurso previamente marcado con `terraform taint`.

13. **`terraform state`**: Gestiona el archivo de estado de Terraform, permitiendo operaciones como mover recursos entre módulos o eliminar recursos del estado.

14. **`terraform graph`**: Genera un gráfico visual de la configuración de Terraform, útil para entender las dependencias entre recursos.



#### 2. **Configurar Credenciales de AWS**: Terraform necesita acceso a tu cuenta de AWS para crear y gestionar recursos. Configura tus credenciales de AWS en `~/.aws/credentials` o usa variables de entorno:

   ```bash
   export AWS_ACCESS_KEY_ID="your_access_key"
   export AWS_SECRET_ACCESS_KEY="your_secret_key"
   ```

#### 3. **Definir la Infraestructura en Terraform**:

   - **Archivo de Configuración**: Crea un archivo `main.tf` para definir los recursos necesarios para tu API y base de datos.

   ```hcl
   provider "aws" {
     region = "us-west-2" // Cambia a tu región preferida
   }

   resource "aws_instance" "api_server" {
     ami           = "ami-0c55b159cbfafe1f0" // Cambia a una AMI adecuada para tu aplicación
     instance_type = "t2.micro"

     tags = {
       Name = "API Server"
     }

     provisioner "remote-exec" {
       inline = [
         "sudo apt-get update",
         "sudo apt-get install -y golang",
         "cd /path/to/your/app",
         "go build -o app",
         "./app"
       ]
     }
   }

   resource "aws_db_instance" "postgresql" {
     allocated_storage    = 20
     engine               = "postgres"
     engine_version       = "13.3"
     instance_class       = "db.t2.micro"
     name                 = "mydb"
     username             = "admin"
     password             = "password"
     parameter_group_name = "default.postgres13"

     skip_final_snapshot = true
   }
   ```

#### 4. **Inicializar Terraform**: Ejecuta `terraform init` en el directorio donde se encuentra tu archivo `main.tf`. Esto descargará los proveedores necesarios.

#### 5. **Planificar la Infraestructura**: Ejecuta `terraform plan` para ver qué cambios se realizarán en tu infraestructura. Esto te permite revisar y confirmar los recursos que se crearán.

#### 6. **Aplicar la Configuración**: Ejecuta `terraform apply` para crear los recursos en AWS. Terraform te pedirá confirmación antes de proceder.

#### 7. **Gestionar Cambios**: Si necesitas cambiar tu infraestructura, edita el archivo `main.tf` y vuelve a ejecutar `terraform plan` y `terraform apply`.

#### 8. **Destruir la Infraestructura**: Si deseas eliminar todos los recursos creados, ejecuta `terraform destroy`.