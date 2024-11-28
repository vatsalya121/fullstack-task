provider "google" {
  credentials = "/Users/vatsalyaparashar/Downloads/rickfav-183b7fc77fa9.json"  # Path to your GCP service account JSON file
  project     = "rickfav"  # Replace with your actual GCP project ID
  region      = "us-central1"
}

# Step 1: Create a persistent disk for MongoDB data storage
resource "google_compute_disk" "mongodb_disk" {
  name  = "mongodb-data-disk"
  size  = 30  # 30 GB free-tier disk
  type  = "pd-standard"  # Standard Persistent Disk
  zone  = "us-central1-a"
}

# Step 2: Define the Compute Engine instance using the f1-micro free-tier
resource "google_compute_instance" "mongodb_instance" {
  name         = "mongodb-instance"
  machine_type = "f1-micro"  # Free-tier machine type
  zone         = "us-central1-a"

  # Boot disk initialization
  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-focal-v20210927"  # Ubuntu OS for the instance
    }
  }

  # Step 3: Attach the persistent disk to the Compute Engine instance
  attached_disk {
    source      = google_compute_disk.mongodb_disk.id  # Attach persistent disk
    device_name = "mongodb-data-disk"
    mode        = "READ_WRITE"
  }

  # Step 4: Add SSH key to the instance for access
  metadata = {
    "ssh-keys" = "ubuntu:${file("/Users/vatsalyaparashar/.ssh/my-gcp-key.pub")}"  # Replace with your public key path
  }

  # Step 5: Allow SSH Access via firewall
  network_interface {
    network = "default"
    access_config {}  # External IP for SSH
  }

  # SSH connection setup
  provisioner "remote-exec" {
    connection {
      type        = "ssh"
      user        = "ubuntu"
      private_key = file("/Users/vatsalyaparashar/.ssh/my-gcp-key")  # Your private key
      host        = self.network_interface[0].access_config[0].nat_ip  # External IP of the instance
    }

    inline = [
      "sudo apt update",
      "sudo apt install -y mongodb",
      "sudo systemctl start mongodb",  # Start MongoDB
      "sudo systemctl enable mongodb"  # Enable MongoDB to start on boot
    ]
  }
}

# Step 6: Firewall rule for SSH access
resource "google_compute_firewall" "mongodb_ssh_access" {
  name    = "allow-ssh"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  source_ranges = ["0.0.0.0/0"]  # Open to all IPs for SSH (can restrict later)
}

# Step 7: Firewall rule for MongoDB access (Restrict to trusted IPs)
resource "google_compute_firewall" "mongodb_access" {
  name    = "allow-mongodb"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["27017"]
  }

  # Replace with your IP range or trusted networks
  source_ranges = ["0.0.0.0/0"]  # Open to all IPs (restrict later)
}

# Step 8: Backup and Restore 
# Use snapshots or custom backup scripts for MongoDB backups.
resource "google_compute_disk_snapshot" "mongodb_backup" {
  name = "mongodb-backup"
  source_disk = google_compute_disk.mongodb_disk.id
  zone         = "us-central1-a"
}

