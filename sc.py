import psycopg2
from faker import Faker
from datetime import datetime, timezone
import random
import sys
import bcrypt

fake = Faker()

DB_PARAMS = {
    'dbname': 'datingapp',
    'user': 'admin',
    'password': 'admin123',
    'host': 'localhost',
    'port': '5432'
}

GENDERS = ['Male', 'Female']
LAT_RANGE = (-11.0, 6.0)
LON_RANGE = (95.0, 141.0)

def create_connection():
    """Create a database connection with PostGIS support"""
    try:
        conn = psycopg2.connect(**DB_PARAMS)
        with conn.cursor() as cur:
            cur.execute("CREATE EXTENSION IF NOT EXISTS postgis;")
            conn.commit()
        return conn
    except psycopg2.OperationalError as e:
        print(f"Could not connect to database: {e}")
        sys.exit(1)

def generate_user_data():
    """Generate random user data"""
    email = fake.email()
    return {
        'created_at': datetime.now(timezone.utc),
        'updated_at': datetime.now(timezone.utc),
        'email': email,
        'password': generate_from_password(email.encode("utf-8"), 10),
        'name': fake.name(),
        'is_premium': random.choice([True, False]),
    }

def generate_profile_data():
    """Generate random profile data with PostGIS point"""
    lat = random.uniform(*LAT_RANGE)
    lon = random.uniform(*LON_RANGE)

    return {
        'created_at': datetime.now(timezone.utc),
        'updated_at': datetime.now(timezone.utc),
        'bio': fake.text(max_nb_chars=300),
        'age': random.randint(18, 65),
        'gender': random.choice(GENDERS),
        'location': f'SRID=4326;POINT({lon} {lat})'
    }

def generate_preference_data():
    """Generate random preference data"""
    return {
        'created_at': datetime.now(timezone.utc),
        'updated_at': datetime.now(timezone.utc),
        'gender_pref': random.choice(GENDERS),
        'distance_pref': random.choice([5, 10, 25, 50, 100])  # km
    }

def generate_from_password(password: bytes, cost: int) -> str:
    if len(password) > 72:
        raise ValueError("Password too long")

    try:
        hashed_password = bcrypt.hashpw(password, bcrypt.gensalt(cost))
        return hashed_password.decode('utf-8')
    except Exception as e:
        raise e



def insert_data(conn, num_records=10):
    """Insert generated data into the database"""
    try:
        with conn.cursor() as cur:
            print(f"Starting to insert {num_records} records...")

            for i in range(num_records):
                # Insert user
                user = generate_user_data()
                cur.execute("""
                    INSERT INTO users
                    (created_at, updated_at, deleted_at, email, password, name, is_premium)
                    VALUES (%s, %s, NULL,%s, %s, %s, %s)
                    RETURNING id
                    """,
                    (
                        user['created_at'], user['updated_at'],
                        user['email'], user['password'],
                        user['name'], user['is_premium']
                    )
                )
                user_id = cur.fetchone()[0]

                profile = generate_profile_data()
                cur.execute("""
                    INSERT INTO profiles
                    (created_at, updated_at, deleted_at, user_id, bio, age, gender, location)
                    VALUES (%s, %s, NULL, %s, %s, %s, %s, ST_GeomFromText(%s))
                    """,
                    (
                        profile['created_at'], profile['updated_at'],
                        user_id, profile['bio'], profile['age'],
                        profile['gender'], profile['location']
                    )
                )

                pref = generate_preference_data()
                cur.execute("""
                    INSERT INTO preferences
                    (created_at, updated_at, deleted_at, user_id, gender_pref, distance_pref)
                    VALUES (%s, %s, NULL, %s, %s, %s)
                    """,
                    (
                        pref['created_at'], pref['updated_at'],
                        user_id, pref['gender_pref'], pref['distance_pref']
                    )
                )

                if (i + 1) % 20 == 0:
                    print(f"Inserted {i + 1} records...")
                    conn.commit()

            conn.commit()
            print("All records inserted successfully!")

    except Exception as e:
        conn.rollback()
        print(f"An error occurred while inserting data: {e}")
        raise

def main():
    conn = None
    try:
        print("Connecting to database...")
        conn = create_connection()
        print("Connected successfully!")

        print("Starting data generation...")
        insert_data(conn, 10)

    except Exception as e:
        print(f"An error occurred: {e}")
        sys.exit(1)
    finally:
        if conn:
            conn.close()
            print("Database connection closed.")

if __name__ == "__main__":
    main()