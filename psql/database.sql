PGDMP                  
    {            postgres    16.0    16.0 (               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    5    postgres    DATABASE     �   CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
    DROP DATABASE postgres;
                postgres    false                       0    0    DATABASE postgres    COMMENT     N   COMMENT ON DATABASE postgres IS 'default administrative connection database';
                   postgres    false    4885                        3079    16384 	   adminpack 	   EXTENSION     A   CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;
    DROP EXTENSION adminpack;
                   false                       0    0    EXTENSION adminpack    COMMENT     M   COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';
                        false    2            �            1259    19326    accounts    TABLE     �   CREATE TABLE public.accounts (
    id integer NOT NULL,
    address character varying(320),
    password character varying(50),
    level integer DEFAULT 1 NOT NULL
);
    DROP TABLE public.accounts;
       public         heap    postgres    false            �            1259    19325    accounts_id_seq    SEQUENCE     �   CREATE SEQUENCE public.accounts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.accounts_id_seq;
       public          postgres    false    221                       0    0    accounts_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.accounts_id_seq OWNED BY public.accounts.id;
          public          postgres    false    220            �            1259    19349    checked_items    TABLE     |   CREATE TABLE public.checked_items (
    person_id integer NOT NULL,
    item_id integer NOT NULL,
    id bigint NOT NULL
);
 !   DROP TABLE public.checked_items;
       public         heap    postgres    false            �            1259    19363    checked_items_id_seq    SEQUENCE     }   CREATE SEQUENCE public.checked_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.checked_items_id_seq;
       public          postgres    false    223                       0    0    checked_items_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.checked_items_id_seq OWNED BY public.checked_items.id;
          public          postgres    false    224            �            1259    19319    items    TABLE     W   CREATE TABLE public.items (
    id integer NOT NULL,
    name character varying(50)
);
    DROP TABLE public.items;
       public         heap    postgres    false            �            1259    19318    items_id_seq    SEQUENCE     �   CREATE SEQUENCE public.items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.items_id_seq;
       public          postgres    false    219                       0    0    items_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.items_id_seq OWNED BY public.items.id;
          public          postgres    false    218            �            1259    19312    people    TABLE     a   CREATE TABLE public.people (
    id integer NOT NULL,
    name character varying(50) NOT NULL
);
    DROP TABLE public.people;
       public         heap    postgres    false            �            1259    19311    people_id_seq    SEQUENCE     �   CREATE SEQUENCE public.people_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.people_id_seq;
       public          postgres    false    217                       0    0    people_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.people_id_seq OWNED BY public.people.id;
          public          postgres    false    216            �            1259    19334    tokens    TABLE     �   CREATE TABLE public.tokens (
    created date DEFAULT CURRENT_DATE,
    updated date DEFAULT CURRENT_DATE,
    expires date DEFAULT (CURRENT_DATE + '01:00:00'::interval),
    token character varying(20) NOT NULL
);
    DROP TABLE public.tokens;
       public         heap    postgres    false            f           2604    19329    accounts id    DEFAULT     j   ALTER TABLE ONLY public.accounts ALTER COLUMN id SET DEFAULT nextval('public.accounts_id_seq'::regclass);
 :   ALTER TABLE public.accounts ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    220    221    221            k           2604    19364    checked_items id    DEFAULT     t   ALTER TABLE ONLY public.checked_items ALTER COLUMN id SET DEFAULT nextval('public.checked_items_id_seq'::regclass);
 ?   ALTER TABLE public.checked_items ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    224    223            e           2604    19322    items id    DEFAULT     d   ALTER TABLE ONLY public.items ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);
 7   ALTER TABLE public.items ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    219    218    219            d           2604    19315 	   people id    DEFAULT     f   ALTER TABLE ONLY public.people ALTER COLUMN id SET DEFAULT nextval('public.people_id_seq'::regclass);
 8   ALTER TABLE public.people ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    217    217                      0    19326    accounts 
   TABLE DATA           @   COPY public.accounts (id, address, password, level) FROM stdin;
    public          postgres    false    221   )                 0    19349    checked_items 
   TABLE DATA           ?   COPY public.checked_items (person_id, item_id, id) FROM stdin;
    public          postgres    false    223   4)       
          0    19319    items 
   TABLE DATA           )   COPY public.items (id, name) FROM stdin;
    public          postgres    false    219   �)                 0    19312    people 
   TABLE DATA           *   COPY public.people (id, name) FROM stdin;
    public          postgres    false    217   ]*                 0    19334    tokens 
   TABLE DATA           B   COPY public.tokens (created, updated, expires, token) FROM stdin;
    public          postgres    false    222   �*                  0    0    accounts_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.accounts_id_seq', 1, false);
          public          postgres    false    220                       0    0    checked_items_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.checked_items_id_seq', 30, true);
          public          postgres    false    224                       0    0    items_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.items_id_seq', 30, true);
          public          postgres    false    218                       0    0    people_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.people_id_seq', 8, true);
          public          postgres    false    216            q           2606    19331    accounts accounts_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.accounts DROP CONSTRAINT accounts_pkey;
       public            postgres    false    221            u           2606    19369     checked_items checked_items_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.checked_items
    ADD CONSTRAINT checked_items_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.checked_items DROP CONSTRAINT checked_items_pkey;
       public            postgres    false    223            o           2606    19324    items items_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.items DROP CONSTRAINT items_pkey;
       public            postgres    false    219            m           2606    19317    people people_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.people
    ADD CONSTRAINT people_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.people DROP CONSTRAINT people_pkey;
       public            postgres    false    217            s           2606    19348    tokens tokens_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_pkey PRIMARY KEY (token);
 <   ALTER TABLE ONLY public.tokens DROP CONSTRAINT tokens_pkey;
       public            postgres    false    222            v           2606    19357 &   checked_items checked_items_item_id_fk    FK CONSTRAINT     �   ALTER TABLE ONLY public.checked_items
    ADD CONSTRAINT checked_items_item_id_fk FOREIGN KEY (item_id) REFERENCES public.items(id);
 P   ALTER TABLE ONLY public.checked_items DROP CONSTRAINT checked_items_item_id_fk;
       public          postgres    false    223    4719    219            w           2606    19352 (   checked_items checked_items_person_id_fk    FK CONSTRAINT     �   ALTER TABLE ONLY public.checked_items
    ADD CONSTRAINT checked_items_person_id_fk FOREIGN KEY (person_id) REFERENCES public.people(id);
 R   ALTER TABLE ONLY public.checked_items DROP CONSTRAINT checked_items_person_id_fk;
       public          postgres    false    223    4717    217                  x������ � �         y   x�ϹD!C�X*f�Y{���1R�% ����u��`r�6VX�a�p�7^48Q#��P�͸����GO��X�J/�8<����8��˦I����v�OW;>A���?��G���      
   �   x�u�;�@ ��b/�n��P�C�	w�GH�8�9�n�젵����������J���p�*�u����ݷ�ɋ	��5w޹sН���t��Io�yj��n�ѐ���3�!�9eȣ�[�!�O���-�F��Ed Z��         I   x�3�tJ��2��L���2�t�/�ML�2�t�K)J-�2��OJ-*�2�t��9�J�,8K�2���b���� we,            x������ � �     