-- 1. Usuarios y Roles
CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    rol ENUM('admin', 'operario') DEFAULT 'operario',
    fecha_registro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 2. Galpones (Infraestructura)
CREATE TABLE galpones (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    capacidad_max INT NOT NULL,
    ubicacion VARCHAR(100)
);

-- 3. Lotes de Aves (Gestión de producción)
CREATE TABLE lotes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    galpon_id INT,
    raza VARCHAR(50),
    cantidad_inicial INT NOT NULL,
    fecha_llegada DATE NOT NULL,
    estado ENUM('activo', 'finalizado') DEFAULT 'activo',
    FOREIGN KEY (galpon_id) REFERENCES galpones(id)
);

-- 4. Registro de Mortalidad
CREATE TABLE mortalidad (
    id INT AUTO_INCREMENT PRIMARY KEY,
    lote_id INT,
    cantidad INT NOT NULL,
    causa VARCHAR(255),
    fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (lote_id) REFERENCES lotes(id) ON DELETE CASCADE
);

-- 5. Sanidad (Vacunas y Tratamientos)
CREATE TABLE sanidad (
    id INT AUTO_INCREMENT PRIMARY KEY,
    lote_id INT,
    tipo EN_ENUM('vacuna', 'medicamento', 'suplemento'),
    descripcion VARCHAR(255),
    fecha_aplicacion DATE,
    FOREIGN KEY (lote_id) REFERENCES lotes(id) ON DELETE CASCADE
);

-- 6. Producción (Peso y Rendimiento)
CREATE TABLE produccion_rendimiento (
    id INT AUTO_INCREMENT PRIMARY KEY,
    lote_id INT,
    peso_promedio_gramos DECIMAL(10,2),
    consumo_alimento_kg DECIMAL(10,2),
    fecha_registro DATE,
    FOREIGN KEY (lote_id) REFERENCES lotes(id) ON DELETE CASCADE
);

-- 7. Contenido Educativo (Tutoriales)
CREATE TABLE tutoriales (
    id INT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(150),
    categoria ENUM('crianza', 'mejora_produccion', 'bioseguridad'),
    contenido TEXT,
    url_video VARCHAR(255),
    fecha_publicacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- --- DATOS DE DEMO ---

INSERT INTO usuarios (nombre, email, password, rol) VALUES 
('Admin Principal', 'admin@granja.com', 'hash_password_seguro', 'admin'),
('Juan Operador', 'juan@granja.com', 'hash_password_seguro', 'operario');

INSERT INTO galpones (nombre, capacidad_max, ubicacion) VALUES 
('Galpón Norte 01', 5000, 'Sector A'),
('Galpón Sur 02', 3000, 'Sector B');

INSERT INTO lotes (galpon_id, raza, cantidad_inicial, fecha_llegada) VALUES 
(1, 'Cobb 500', 4500, '2026-01-10'),
(2, 'Ross 308', 2800, '2026-02-01');

INSERT INTO mortalidad (lote_id, cantidad, causa) VALUES 
(1, 5, 'Estrés por transporte'),
(1, 2, 'Causas naturales');

INSERT INTO tutoriales (titulo, categoria, contenido) VALUES 
('Manejo de la primera semana', 'crianza', 'La temperatura debe oscilar entre 32-34°C...'),
('Optimización del FCR', 'mejora_produccion', 'Cómo mejorar el índice de conversión alimenticia...');
